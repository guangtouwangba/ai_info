# Large Language Models Empowered Personalized Web Agents

**URL**: http://arxiv.org/abs/2410.17236v2

## 原始摘要

Web agents have emerged as a promising direction to automate Web task
completion based on user instructions, significantly enhancing user experience.
Recently, Web agents have evolved from traditional agents to Large Language
Models (LLMs)-based Web agents. Despite their success, existing LLM-based Web
agents overlook the importance of personalized data (e.g., user profiles and
historical Web behaviors) in assisting the understanding of users' personalized
instructions and executing customized actions. To overcome the limitation, we
first formulate the task of LLM-empowered personalized Web agents, which
integrate personalized data and user instructions to personalize instruction
comprehension and action execution. To address the absence of a comprehensive
evaluation benchmark, we construct a Personalized Web Agent Benchmark
(PersonalWAB), featuring user instructions, personalized user data, Web
functions, and two evaluation paradigms across three personalized Web tasks.
Moreover, we propose a Personalized User Memory-enhanced Alignment (PUMA)
framework to adapt LLMs to the personalized Web agent task. PUMA utilizes a
memory bank with a task-specific retrieval strategy to filter relevant
historical Web behaviors. Based on the behaviors, PUMA then aligns LLMs for
personalized action execution through fine-tuning and direct preference
optimization. Extensive experiments validate the superiority of PUMA over
existing Web agents on PersonalWAB.


## AI 摘要

这篇论文提出了一种基于大语言模型（LLM）的个性化网络代理（Personalized Web Agent）任务，旨在利用用户个性化数据（如个人资料和历史行为）来更好地理解和执行用户指令。为解决缺乏评估基准的问题，研究者构建了PersonalWAB基准，包含用户指令、个性化数据、网络功能和两种评估范式。此外，论文提出了PUMA框架，通过记忆库检索相关历史行为，并利用微调和偏好优化使LLM适应个性化任务。实验表明，PUMA在PersonalWAB上优于现有网络代理。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-25T23:02:17Z
- **目录日期**: 2025-03-25
