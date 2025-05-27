# Evaluating Large Language Models for Code Review

**URL**: http://arxiv.org/abs/2505.20206v1

## 原始摘要

Context: Code reviews are crucial for software quality. Recent AI advances
have allowed large language models (LLMs) to review and fix code; now, there
are tools that perform these reviews. However, their reliability and accuracy
have not yet been systematically evaluated. Objective: This study compares
different LLMs' performance in detecting code correctness and suggesting
improvements. Method: We tested GPT4o and Gemini 2.0 Flash on 492 AI generated
code blocks of varying correctness, along with 164 canonical code blocks from
the HumanEval benchmark. To simulate the code review task objectively, we
expected LLMs to assess code correctness and improve the code if needed. We ran
experiments with different configurations and reported on the results. Results:
With problem descriptions, GPT4o and Gemini 2.0 Flash correctly classified code
correctness 68.50% and 63.89% of the time, respectively, and corrected the code
67.83% and 54.26% of the time for the 492 code blocks of varying correctness.
Without problem descriptions, performance declined. The results for the 164
canonical code blocks differed, suggesting that performance depends on the type
of code. Conclusion: LLM code reviews can help suggest improvements and assess
correctness, but there is a risk of faulty outputs. We propose a process that
involves humans, called the "Human in the loop LLM Code Review" to promote
knowledge sharing while mitigating the risk of faulty outputs.


## AI 摘要

该研究评估了GPT4o和Gemini 2.0 Flash在代码审查中的表现。测试使用492个AI生成代码块和164个HumanEval基准代码块，要求模型判断正确性并提出改进。结果显示，在有问题描述时，GPT4o和Gemini正确分类代码的准确率分别为68.50%和63.89%，修正代码的成功率为67.83%和54.26%；无描述时性能下降。不同代码类型的表现存在差异。研究表明，虽然LLM能辅助代码审查，但存在错误输出风险，因此建议采用"Human in the loop LLM Code Review"流程，结合人工干预以确保质量并促进知识共享。（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T09:01:41Z
- **目录日期**: 2025-05-27
