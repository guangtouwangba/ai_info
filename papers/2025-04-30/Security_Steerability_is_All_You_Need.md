# Security Steerability is All You Need

**URL**: http://arxiv.org/abs/2504.19521v2

## 原始摘要

The adoption of Generative AI (GenAI) in various applications inevitably
comes with expanding the attack surface, combining new security threats along
with the traditional ones. Consequently, numerous research and industrial
initiatives aim to mitigate these security threats in GenAI by developing
metrics and designing defenses. However, while most of the GenAI security work
focuses on universal threats (e.g. manipulating the LLM to generate forbidden
content), there is significantly less discussion on application-level security
and how to mitigate it. Thus, in this work we adopt an application-centric
approach to GenAI security, and show that while LLMs cannot protect against
ad-hoc application specific threats, they can provide the framework for
applications to protect themselves against such threats. Our first contribution
is defining Security Steerability - a novel security measure for LLMs,
assessing the model's capability to adhere to strict guardrails that are
defined in the system prompt ('Refrain from discussing about politics'). These
guardrails, in case effective, can stop threats in the presence of malicious
users who attempt to circumvent the application and cause harm to its
providers. Our second contribution is a methodology to measure the security
steerability of LLMs, utilizing two newly-developed datasets: VeganRibs
assesses the LLM behavior in forcing specific guardrails that are not security
per se in the presence of malicious user that uses attack boosters (jailbreaks
and perturbations), and ReverseText takes this approach further and measures
the LLM ability to force specific treatment of the user input as plain text
while do user try to give it additional meanings...


## AI 摘要

本文提出了一种以应用为中心的生成式AI安全方法，强调应用级安全防护的重要性。研究定义了"安全可控性"这一新指标，评估大语言模型(LLM)遵循系统提示中严格防护规则的能力(如"避免讨论政治")。作者开发了两种测试方法：VeganRibs检测模型在面对恶意攻击(越狱和干扰)时执行非安全类规则的能力；ReverseText则测试模型将用户输入视为纯文本处理的能力。研究表明，虽然LLM无法完全防范特定应用威胁，但能为应用提供自我保护框架，这对防范恶意用户绕过系统造成危害具有重要意义。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T09:02:17Z
- **目录日期**: 2025-04-30
