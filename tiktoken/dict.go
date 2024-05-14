package tiktoken

import _ "embed"

var (
    //go:embed qwen.tiktoken
    QwenEmbedDict []byte
)
