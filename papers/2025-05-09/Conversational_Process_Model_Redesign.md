# Conversational Process Model Redesign

**URL**: http://arxiv.org/abs/2505.05453v1

## 原始摘要

With the recent success of large language models (LLMs), the idea of
AI-augmented Business Process Management systems is becoming more feasible. One
of their essential characteristics is the ability to be conversationally
actionable, allowing humans to interact with the LLM effectively to perform
crucial process life cycle tasks such as process model design and redesign.
However, most current research focuses on single-prompt execution and
evaluation of results, rather than on continuous interaction between the user
and the LLM. In this work, we aim to explore the feasibility of using LLMs to
empower domain experts in the creation and redesign of process models in an
iterative and effective way. The proposed conversational process model redesign
(CPD) approach receives as input a process model and a redesign request by the
user in natural language. Instead of just letting the LLM make changes, the LLM
is employed to (a) identify process change patterns from literature, (b)
re-phrase the change request to be aligned with an expected wording for the
identified pattern (i.e., the meaning), and then to (c) apply the meaning of
the change to the process model. This multi-step approach allows for
explainable and reproducible changes. In order to ensure the feasibility of the
CPD approach, and to find out how well the patterns from literature can be
handled by the LLM, we performed an extensive evaluation. The results show that
some patterns are hard to understand by LLMs and by users. Within the scope of
the study, we demonstrated that users need support to describe the changes
clearly. Overall the evaluation shows that the LLMs can handle most changes
well according to a set of completeness and correctness criteria.


## AI 摘要

随着大语言模型(LLMs)的成功，AI增强的业务流程管理系统变得可行。关键特性是支持对话式交互，使用户能通过自然语言与LLM协作完成流程模型设计与重构。本研究提出了一种对话式流程模型重构(CPD)方法，通过多步骤实现可解释的修改：(1)识别流程变更模式；(2)调整变更请求的表述；(3)将变更应用到模型。评估表明，LLM能较好地处理多数变更模式，但某些模式对LLM和用户仍具挑战性。用户需要支持以清晰描述变更需求。总体而言，LLM在完整性和正确性标准下表现良好。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T23:01:10Z
- **目录日期**: 2025-05-09
