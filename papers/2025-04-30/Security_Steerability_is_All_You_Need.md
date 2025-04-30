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

本文提出了一种以应用为中心的生成式AI（GenAI）安全方法，强调当前研究多关注通用威胁（如生成违禁内容），而忽视了应用级安全问题。作者定义了"安全可控性"（Security Steerability）这一新安全指标，评估大语言模型（LLM）遵循系统提示中严格防护规则（如"不讨论政治"）的能力，以抵御恶意用户绕过应用的攻击。研究还开发了两种新数据集（VeganRibs和ReverseText）来衡量LLM的安全可控性，前者测试模型在攻击增强（越狱和干扰）下执行非安全防护规则的能力，后者评估模型将用户输入视为纯文本处理的准确性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T21:01:57Z
- **目录日期**: 2025-04-30
