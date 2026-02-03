# alidns

阿里云 DNS CLI，支持以下子命令：

- `add`: 添加 DNS 记录
- `del`: 删除 DNS 记录
- `query`: 查询 DNS 记录
- `update`: 修改 DNS 记录

## 用途与输出

- 统一入口：通过同一套参数风格调用 add/del/query/update。
- 输出格式：支持 `--output json|pretty`，默认 `pretty`。
- 子命令上的 `--output` 优先级高于全局 `--output`。

## 构建与运行

### 依赖

- Go 1.21+

### 本地构建

```bash
go build -ldflags "-s -w" -o alidns ./cmd/alidns
```

### 交叉编译示例

```bash
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o alidns.exe ./cmd/alidns
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o alidns ./cmd/alidns
```

### 运行帮助

```bash
./alidns -h
./alidns help add
./alidns add -h
```

## 全局参数

- `--output string`：输出格式，`json|pretty`，默认 `pretty`
- `-h, --help`：显示帮助

## 子命令详解

### add

新增解析记录。

```bash
alidns add -ak AK -sk SK -domain example.com -name www -type A -value 1.2.3.4 \
  [--ttl 600] [--priority 1] [--line default] [--output json|pretty]
```

参数：
- 必填：`-ak`、`-sk`、`-domain`、`-name`、`-type`、`-value`
- 可选：`-ttl`（默认 `600`）、`-priority`（默认 `1`）、`-line`（默认 `default`）、`--output`

示例：

```bash
alidns add -ak AK -sk SK -domain example.com -name www -type A -value 1.2.3.4
alidns add -ak AK -sk SK -domain example.com -name @ -type TXT -value hello --output json
```

### del

按主域名 + 主机记录 + 记录类型删除解析记录。

```bash
alidns del -ak AK -sk SK -domain example.com -name www -type A [--output json|pretty]
```

参数：
- 必填：`-ak`、`-sk`、`-domain`、`-name`、`-type`
- 可选：`--output`

示例：

```bash
alidns del -ak AK -sk SK -domain example.com -name www -type A
```

### query

查询主域名下的记录列表。

```bash
alidns query -ak AK -sk SK -domain example.com [--output json|pretty]
```

参数：
- 必填：`-ak`、`-sk`、`-domain`
- 可选：`--output`

说明：
- 只输出 Record 列表。
- 无记录时输出 `[]`。

示例：

```bash
alidns query -ak AK -sk SK -domain example.com
alidns query -ak AK -sk SK -domain example.com --output json
```

### update

按记录 ID 修改记录内容。

```bash
alidns update -ak AK -sk SK -id RECORD_ID -name www -type A -value 1.2.3.4 \
  [--ttl 600] [--priority 1] [--line default] [--output json|pretty]
```

参数：
- 必填：`-ak`、`-sk`、`-id`、`-name`、`-type`、`-value`
- 可选：`-ttl`（默认 `600`）、`-priority`（默认 `1`）、`-line`（默认 `default`）、`--output`

示例：

```bash
alidns update -ak AK -sk SK -id RECORD_ID -name www -type A -value 1.2.3.4
```

## 输出格式

- `--output pretty`：多行缩进 JSON，便于人工阅读。
- `--output json`：单行 JSON，便于脚本处理。

```bash
alidns query -ak AK -sk SK -domain example.com --output json | jq .
```

## 开发与测试

```bash
go test ./...
```
