# Gaming Tool Preferences in Agentic LLMs

**URL**: http://arxiv.org/abs/2505.18135v1

## 原始摘要

Large language models (LLMs) can now access a wide range of external tools,
thanks to the Model Context Protocol (MCP). This greatly expands their
abilities as various agents. However, LLMs rely entirely on the text
descriptions of tools to decide which ones to use--a process that is
surprisingly fragile. In this work, we expose a vulnerability in prevalent
tool/function-calling protocols by investigating a series of edits to tool
descriptions, some of which can drastically increase a tool's usage from LLMs
when competing with alternatives. Through controlled experiments, we show that
tools with properly edited descriptions receive over 10 times more usage from
GPT-4.1 and Qwen2.5-7B than tools with original descriptions. We further
evaluate how various edits to tool descriptions perform when competing directly
with one another and how these trends generalize or differ across a broader set
of 10 different models. These phenomenons, while giving developers a powerful
way to promote their tools, underscore the need for a more reliable foundation
for agentic LLMs to select and utilize tools and resources.


## AI 摘要

研究发现大语言模型(LLMs)在调用外部工具时存在严重漏洞：仅依赖工具文本描述进行选择的过程非常脆弱。通过修改工具描述，某些工具的使用率可提升10倍以上(GPT-4.1和Qwen2-7B测试)。实验表明，优化后的描述能显著增加工具被选中的概率，不同模型间表现趋势相似。这一现象虽为开发者推广工具提供了捷径，但也凸显出现有代理协议在选择工具时缺乏可靠机制，亟需建立更稳健的工具调用基础架构。研究测试了10种不同模型，结果均显示这一脆弱性普遍存在。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T11:01:31Z
- **目录日期**: 2025-05-26
