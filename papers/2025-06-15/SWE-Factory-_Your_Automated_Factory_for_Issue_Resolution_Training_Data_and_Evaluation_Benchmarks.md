# SWE-Factory: Your Automated Factory for Issue Resolution Training Data and Evaluation Benchmarks

**URL**: http://arxiv.org/abs/2506.10954v1

## 原始摘要

Constructing large-scale datasets for the GitHub issue resolution task is
crucial for both training and evaluating the software engineering capabilities
of Large Language Models (LLMs). However, the traditional process for creating
such benchmarks is notoriously challenging and labor-intensive, particularly in
the stages of setting up evaluation environments, grading test outcomes, and
validating task instances. In this paper, we propose SWE-Factory, an automated
pipeline designed to address these challenges. To tackle these issues, our
pipeline integrates three core automated components. First, we introduce
SWE-Builder, a multi-agent system that automates evaluation environment
construction, which employs four specialized agents that work in a
collaborative, iterative loop and leverages an environment memory pool to
enhance efficiency. Second, we introduce a standardized, exit-code-based
grading method that eliminates the need for manually writing custom parsers.
Finally, we automate the fail2pass validation process using these reliable exit
code signals. Experiments on 671 issues across four programming languages show
that our pipeline can effectively construct valid task instances; for example,
with GPT-4.1-mini, our SWE-Builder constructs 269 valid instances at $0.045 per
instance, while with Gemini-2.5-flash, it achieves comparable performance at
the lowest cost of $0.024 per instance. We also demonstrate that our
exit-code-based grading achieves 100% accuracy compared to manual inspection,
and our automated fail2pass validation reaches a precision of 0.92 and a recall
of 1.00. We hope our automated pipeline will accelerate the collection of
large-scale, high-quality GitHub issue resolution datasets for both training
and evaluation. Our code and datasets are released at
https://github.com/DeepSoftwareAnalytics/swe-factory.


## AI 摘要

构建大规模GitHub问题解决数据集对训练和评估大语言模型（LLM）的软件工程能力至关重要，但传统方法在搭建评估环境、测试评分和验证任务实例方面耗时费力。为此，研究者提出SWE-Factory自动化流水线，整合三个核心组件：1）多智能体系统SWE-Builder自动构建评估环境；2）基于退出码的标准化评分方法，无需人工编写解析器；3）利用可靠退出码信号自动化fail2pass验证。实验表明，该流水线能高效构建有效任务实例（如GPT-4.1-mini以0.045美元/实例生成269个实例），评分准确率达100%，验证精确度0.92，召回率1.00。代码与数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-15T21:01:25Z
- **目录日期**: 2025-06-15
