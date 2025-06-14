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

这篇论文提出了TDD-Bench-Verified基准和基于LLM的Otter系统，用于从问题描述生成测试用例，支持测试驱动开发(TDD)和验证软件工程代理生成的代码补丁。Otter通过规则分析检查和修复LLM输出，并采用自反思动作规划器。实验表明，Otter在从问题生成测试方面优于现有系统，同时也能提升生成补丁系统的表现。该研究旨在帮助开发者更高效解决问题并产生更健壮、经过充分测试的代码，填补了现有研究中从问题直接生成测试的空白。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T23:02:11Z
- **目录日期**: 2025-06-02
