# PandaGuard: Systematic Evaluation of LLM Safety against Jailbreaking Attacks

**URL**: http://arxiv.org/abs/2505.13862v3

## 原始摘要

Large language models (LLMs) have achieved remarkable capabilities but remain
vulnerable to adversarial prompts known as jailbreaks, which can bypass safety
alignment and elicit harmful outputs. Despite growing efforts in LLM safety
research, existing evaluations are often fragmented, focused on isolated attack
or defense techniques, and lack systematic, reproducible analysis. In this
work, we introduce PandaGuard, a unified and modular framework that models LLM
jailbreak safety as a multi-agent system comprising attackers, defenders, and
judges. Our framework implements 19 attack methods and 12 defense mechanisms,
along with multiple judgment strategies, all within a flexible plugin
architecture supporting diverse LLM interfaces, multiple interaction modes, and
configuration-driven experimentation that enhances reproducibility and
practical deployment. Built on this framework, we develop PandaBench, a
comprehensive benchmark that evaluates the interactions between these
attack/defense methods across 49 LLMs and various judgment approaches,
requiring over 3 billion tokens to execute. Our extensive evaluation reveals
key insights into model vulnerabilities, defense cost-performance trade-offs,
and judge consistency. We find that no single defense is optimal across all
dimensions and that judge disagreement introduces nontrivial variance in safety
assessments. We release the code, configurations, and evaluation results to
support transparent and reproducible research in LLM safety.


## AI 摘要

本文介绍了PandaGuard框架和PandaBench基准测试，用于系统评估大语言模型(LLMs)的对抗性攻击防御能力。该框架采用多代理系统(攻击者、防御者和评判者)，整合了19种攻击方法、12种防御机制和多种评判策略，支持灵活配置和可重复实验。基于此开发的PandaBench测试了49种LLM在攻击/防御交互中的表现，消耗超30亿token。研究发现：没有单一防御方法能在所有维度表现最优；评判标准差异会显著影响安全评估结果。研究提供了代码、配置和评估结果以促进LLM安全研究的透明性和可重复性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T17:02:51Z
- **目录日期**: 2025-05-27
