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

这篇论文提出了PandaGuard框架和PandaBench基准，用于系统评估大语言模型(LLM)的越狱安全风险。该框架采用多智能体架构(攻击者、防御者和评判者)，整合了19种攻击方法、12种防御机制和多种评判策略，支持可配置的实验设置。基于此构建的PandaBench基准测试了49个LLM模型，消耗超过30亿token。研究发现：没有单一防御能应对所有攻击；评判标准差异会导致安全评估显著波动；防御措施存在成本-性能权衡。研究团队开源了代码和配置以促进可复现的LLM安全研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T22:02:24Z
- **目录日期**: 2025-05-27
