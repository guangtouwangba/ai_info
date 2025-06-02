# Otter: Generating Tests from Issues to Validate SWE Patches

**URL**: http://arxiv.org/abs/2502.05368v2

## 原始摘要

While there has been plenty of work on generating tests from existing code,
there has been limited work on generating tests from issues. A correct test
must validate the code patch that resolves the issue. This paper focuses on the
scenario where that code patch does not yet exist. Doing so supports two major
use-cases. First, it supports TDD (test-driven development), the discipline of
"test first, write code later" that has well-documented benefits for human
software engineers. Second, it also validates SWE (software engineering)
agents, which generate code patches for resolving issues. This paper introduces
TDD-Bench-Verified, a benchmark for generating tests from issues, and Otter, an
LLM-based solution for this task. Otter augments LLMs with rule-based analysis
to check and repair their outputs, and introduces a novel self-reflective
action planner. Experiments show Otter outperforming state-of-the-art systems
for generating tests from issues, in addition to enhancing systems that
generate patches from issues. We hope that Otter helps make developers more
productive at resolving issues and leads to more robust, well-tested code.


## AI 摘要

这篇论文提出了TDD-Bench-Verified基准测试和基于LLM的解决方案Otter，用于从问题描述生成测试用例，支持测试驱动开发(TDD)和验证软件工程代理生成的代码补丁。Otter通过规则分析检查和修复LLM输出，并引入创新的自反式行动规划器。实验表明，Otter在从问题生成测试方面优于现有技术，并能提升生成补丁系统的性能。该研究旨在提高开发者解决问题的效率，促进更健壮、经过充分测试的代码开发。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T04:08:32Z
- **目录日期**: 2025-06-02
