# Can We Trust Embodied Agents? Exploring Backdoor Attacks against Embodied LLM-based Decision-Making Systems

**URL**: http://arxiv.org/abs/2405.20774v3

## 原始摘要

Large Language Models (LLMs) have shown significant promise in real-world
decision-making tasks for embodied artificial intelligence, especially when
fine-tuned to leverage their inherent common sense and reasoning abilities
while being tailored to specific applications. However, this fine-tuning
process introduces considerable safety and security vulnerabilities, especially
in safety-critical cyber-physical systems. In this work, we propose the first
comprehensive framework for Backdoor Attacks against LLM-based Decision-making
systems (BALD) in embodied AI, systematically exploring the attack surfaces and
trigger mechanisms. Specifically, we propose three distinct attack mechanisms:
word injection, scenario manipulation, and knowledge injection, targeting
various components in the LLM-based decision-making pipeline. We perform
extensive experiments on representative LLMs (GPT-3.5, LLaMA2, PaLM2) in
autonomous driving and home robot tasks, demonstrating the effectiveness and
stealthiness of our backdoor triggers across various attack channels, with
cases like vehicles accelerating toward obstacles and robots placing knives on
beds. Our word and knowledge injection attacks achieve nearly 100% success rate
across multiple models and datasets while requiring only limited access to the
system. Our scenario manipulation attack yields success rates exceeding 65%,
reaching up to 90%, and does not require any runtime system intrusion. We also
assess the robustness of these attacks against defenses, revealing their
resilience. Our findings highlight critical security vulnerabilities in
embodied LLM systems and emphasize the urgent need for safeguarding these
systems to mitigate potential risks.


## AI 摘要

本研究提出了首个针对基于大语言模型(LLM)的具身AI决策系统的后门攻击框架BALD，系统探索了攻击面和触发机制。通过词注入、场景操纵和知识注入三种攻击方式，在自动驾驶和家用机器人任务中对GPT-3.5等主流模型进行测试，攻击成功率最高达100%。攻击可导致车辆加速撞向障碍物或机器人将刀具放在床上等危险行为，且多数攻击无需系统入侵。实验表明这些攻击对现有防御措施具有强鲁棒性，揭示了具身LLM系统的严重安全隐患，亟需开发有效的安全防护机制。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-02T02:30:13Z
- **目录日期**: 2025-05-02
