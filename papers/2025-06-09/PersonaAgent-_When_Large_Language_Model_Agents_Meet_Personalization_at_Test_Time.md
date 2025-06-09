# PersonaAgent: When Large Language Model Agents Meet Personalization at Test Time

**URL**: http://arxiv.org/abs/2506.06254v1

## 原始摘要

Large Language Model (LLM) empowered agents have recently emerged as advanced
paradigms that exhibit impressive capabilities in a wide range of domains and
tasks. Despite their potential, current LLM agents often adopt a
one-size-fits-all approach, lacking the flexibility to respond to users'
varying needs and preferences. This limitation motivates us to develop
PersonaAgent, the first personalized LLM agent framework designed to address
versatile personalization tasks. Specifically, PersonaAgent integrates two
complementary components - a personalized memory module that includes episodic
and semantic memory mechanisms; a personalized action module that enables the
agent to perform tool actions tailored to the user. At the core, the persona
(defined as unique system prompt for each user) functions as an intermediary:
it leverages insights from personalized memory to control agent actions, while
the outcomes of these actions in turn refine the memory. Based on the
framework, we propose a test-time user-preference alignment strategy that
simulate the latest n interactions to optimize the persona prompt, ensuring
real-time user preference alignment through textual loss feedback between
simulated and ground-truth responses. Experimental evaluations demonstrate that
PersonaAgent significantly outperforms other baseline methods by not only
personalizing the action space effectively but also scaling during test-time
real-world applications. These results underscore the feasibility and potential
of our approach in delivering tailored, dynamic user experiences.


## AI 摘要

PersonaAgent是一个新型的个性化大语言模型（LLM）代理框架，通过整合个性化记忆模块（情景和语义记忆）和个性化行动模块（定制化工具操作），动态响应用户需求。其核心是用户独特的"角色"系统提示，协调记忆与行动间的交互，并通过测试时用户偏好对齐策略实时优化提示。实验表明，PersonaAgent在个性化行动空间和实际应用扩展性上显著优于基线方法，展现了提供动态定制用户体验的潜力。该框架突破了传统LLM代理"一刀切"的局限，首次实现了多功能个性化任务处理。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-09T10:00:54Z
- **目录日期**: 2025-06-09
