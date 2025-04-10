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

FEABench是一个评估大语言模型(LLMs)和LLM代理使用有限元分析(FEA)解决物理、数学和工程问题能力的基准。它通过自然语言问题描述和操作COMSOL Multiphysics软件进行端到端求解，并设计了能通过API与软件交互、检查输出并迭代改进解决方案的智能代理。最佳策略可生成88%可执行API调用。这种能力将增强LLMs的数值求解精度，推动工程自动化发展。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T01:28:23Z
- **目录日期**: 2025-04-10
