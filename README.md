# Artifact Provenance Verifier

纯Go企业级软件供应链证明与SBOM验证服务。服务接收制品摘要、in-toto/DSSE风格证明、SPDX或CycloneDX SBOM和漏洞情报，执行可追踪的发布门禁验证。

## 快速启动

```sh
API_KEY=dev-secret go run ./cmd/server
```

默认监听`:8092`。除`GET`健康和指标接口外，API请求需要`X-API-Key: dev-secret`。

## 端到端示例

```sh
curl -H 'X-API-Key: dev-secret' -H 'Content-Type: application/json' -d '{"organization":"acme","repository":"payments","name":"api","digest":"sha256:0123456789abcdef0123456789abcdef","algorithm":"sha256","architecture":"amd64","version":"1.2.3"}' http://localhost:8092/v1/artifacts
curl -H 'X-API-Key: dev-secret' -H 'Content-Type: application/json' -d '{"subject":"sha256:0123456789abcdef0123456789abcdef","predicate_type":"https://slsa.dev/provenance/v1","body_digest":"sha256:abcdef0123456789abcdef0123456789","signature":"demo"}' http://localhost:8092/v1/artifacts/art-*/attestations
```

实际请求中将上一步返回的制品`id`替换到路径。SBOM使用`format`为`SPDX`或`CycloneDX`，组件至少包含`name`和`version`。验证请求示例：

```json
{"artifact_id":"art-...","require_attestation":true,"require_sbom":true,"max_severity":"high"}
```

## 架构

`internal`按artifact、attestation、sbom、dependency、vulnerability、verification、trust、storage、webhook和worker领域拆分；每个领域提供domain/application/adapter/infrastructure层。默认使用线程安全内存仓储，`internal/storage/adapter`保留PostgreSQL健康检查接口；迁移位于`migrations/001_initial.sql`。所有服务使用context、包装错误、构造函数注入和结构化`slog`日志。

## 运维

提供`/healthz`、`/readyz`和Prometheus文本格式`/metrics`，支持请求体上限、超时、API Key认证、优雅停机。生产环境可使用`deploy/Dockerfile`和`deploy/docker-compose.yml`。
