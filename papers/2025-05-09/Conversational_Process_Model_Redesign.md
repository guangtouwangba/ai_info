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

大型语言模型(LLMs)为AI增强的业务流程管理系统提供了新可能。本研究提出了一种对话式流程模型重构(CPD)方法，通过多步骤交互支持用户自然语言修改请求：首先识别流程变更模式，然后调整请求表述以匹配模式术语，最后应用变更。这种方法确保了修改的可解释性和可重复性。评估表明，虽然某些模式对LLM和用户较难理解，且用户需要辅助来清晰描述变更，但LLM能较好地处理大多数变更请求，满足完整性和正确性标准。研究突出了持续交互在流程模型迭代改进中的重要性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T11:01:15Z
- **目录日期**: 2025-05-09
