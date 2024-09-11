package llm

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"os"
	"search/internal/schema"
	"strings"
)

type Enhancer struct{}

var (
	systemMessageContent = "You are a helpful assistant that create final answer for given text.\n Your main task is providing enriched result for given query.\n The query will be provided. Alongside the query list of web results will be provided in this format Title, Snippet, Url.\n You can use the Snippet and Title to find matched results and if so, you use the Url to get much information from that website.\nResult should be well-crafted.\nYou can harmonize the web results and your internal knowledge.\nFinally, add link to the result which you sourced."
	userMessageContent   = "You are a helpful assistant. Please analyze the following user query and create answer for it, the main query is : %s.\nIn the web results, you can have small text snippet, title and url for that.\n If you think that any snippet is related to question go to provided link to fetch much more information to use in your final answer.\nFor that question here is the some web results you can use:\n[\n%s\n]"
)

func (e Enhancer) messageBuilder(query string, results schema.EngineResponseList) []openai.ChatCompletionMessage {
	systemMessage := openai.ChatCompletionMessage{Role: "system", Content: systemMessageContent}

	webResultsList := make([]string, 0, len(results))
	for _, result := range results {
		webResults := fmt.Sprintf("Title:%s\nSnipet:%s\nUrl:%s\n\n", result.Title, result.Snippet, result.Url)
		webResultsList = append(webResultsList, webResults)
	}
	webResultsString := strings.Join(webResultsList, "")
	userMessageContent = fmt.Sprintf(userMessageContent, query, webResultsString)
	userMessage := openai.ChatCompletionMessage{Role: "user", Content: userMessageContent}

	messages := []openai.ChatCompletionMessage{systemMessage, userMessage}
	return messages

}

func (e Enhancer) FinalizeResult(query string, result schema.EngineResponseList) (string, error) {

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	ctx := context.Background()
	messages := e.messageBuilder(query, result)
	completion := openai.ChatCompletionRequest{Model: openai.GPT4o20240806, Messages: messages}
	res, err := client.CreateChatCompletion(ctx, completion)
	if err != nil {
		return "", err
	}

	content := res.Choices[0].Message.Content

	return content, nil
}
