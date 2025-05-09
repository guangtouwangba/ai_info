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

随着大语言模型(LLMs)的发展，AI增强的业务流程管理系统变得可行。当前研究多关注单次提示执行，而本文提出"对话式流程模型重构"(CPD)方法，通过多步骤交互支持用户迭代修改流程模型：LLM先识别流程变更模式，调整用户自然语言请求以匹配模式术语，再应用变更。评估表明，LLMs能较好处理多数变更模式，但某些模式对LLM和用户仍具挑战性，用户需要辅助以清晰描述变更需求。该方法实现了可解释、可复现的流程修改，验证了LLMs在业务流程管理中的实用性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T15:01:17Z
- **目录日期**: 2025-05-09
