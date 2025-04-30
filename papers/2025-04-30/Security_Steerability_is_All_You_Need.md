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

这篇论文探讨了生成式AI（GenAI）应用中的安全问题，提出了一种以应用为中心的安全防护方法。研究重点在于定义并衡量"安全可控性"（Security Steerability），即大语言模型（LLM）遵循系统提示中严格护栏规则（如"避免讨论政治"）的能力，以防止恶意用户绕过应用限制。作者开发了两种新数据集（VeganRibs和ReverseText）来测试LLM在遭遇攻击（如越狱和干扰）时的防护表现，评估其能否强制处理用户输入为纯文本而忽略附加恶意意图。该研究为应用级GenAI安全防护提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T11:02:06Z
- **目录日期**: 2025-04-30
