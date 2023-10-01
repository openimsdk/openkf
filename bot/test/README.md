# Test scripts

## Test Milvus standalone startup

### 1. Start Milvus standalone

```bash
# DIR: /openkf/deploy/docker-compose
docker-compose up -d
```

### 2. Check with scripts

```bash
# DIR: /openkf/bot/test
# pymilvus==2.1.3

python3 test_milvus.py
```

### 3. Outputs

```bash

=== start connecting to Milvus     ===

Does collection hello_milvus exist in Milvus: False

=== Create collection `hello_milvus` ===


=== Start inserting entities       ===

Number of entities in Milvus: 3000

=== Start Creating index IVF_FLAT  ===


=== Start loading                  ===


=== Start searching based on vector similarity ===
....
```

## Test langchain with llms

### 1. Start LLM server

```bash
# DIR: /openkf/bot/llm

python3 fastchat_api_runner.py
```

### 2. Setup environment

```bash
export OPENAI_API_BASE=http://localhost:8888/v1
export OPENAI_API_KEY=EMPTY
```

### 3. Test chats

```bash
# DIR: /openkf/bot/test
python3 test_langchain.py
```