# FEABench: Evaluating Language Models on Multiphysics Reasoning Ability

**URL**: http://arxiv.org/abs/2504.06260v1

## 原始摘要

Building precise simulations of the real world and invoking numerical solvers
to answer quantitative problems is an essential requirement in engineering and
science. We present FEABench, a benchmark to evaluate the ability of large
language models (LLMs) and LLM agents to simulate and solve physics,
mathematics and engineering problems using finite element analysis (FEA). We
introduce a comprehensive evaluation scheme to investigate the ability of LLMs
to solve these problems end-to-end by reasoning over natural language problem
descriptions and operating COMSOL Multiphysics$^\circledR$, an FEA software, to
compute the answers. We additionally design a language model agent equipped
with the ability to interact with the software through its Application
Programming Interface (API), examine its outputs and use tools to improve its
solutions over multiple iterations. Our best performing strategy generates
executable API calls 88% of the time. LLMs that can successfully interact with
and operate FEA software to solve problems such as those in our benchmark would
push the frontiers of automation in engineering. Acquiring this capability
would augment LLMs' reasoning skills with the precision of numerical solvers
and advance the development of autonomous systems that can tackle complex
problems in the real world. The code is available at
https://github.com/google/feabench


## AI 摘要

FEABench是一个评估大语言模型(LLMs)和LLM代理使用有限元分析(FEA)解决物理、数学和工程问题能力的基准。该研究开发了综合评估方案，测试LLMs通过自然语言理解问题描述并操作COMSOL Multiphysics软件进行端到端求解的能力。研究者还设计了能通过API与软件交互的智能代理，可检查输出并通过多次迭代改进解决方案。最佳策略能生成88%可执行的API调用。这种结合LLM推理与数值求解器精度的能力将推动工程自动化发展，使自主系统能解决现实世界复杂问题。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T04:00:58Z
- **目录日期**: 2025-04-09
