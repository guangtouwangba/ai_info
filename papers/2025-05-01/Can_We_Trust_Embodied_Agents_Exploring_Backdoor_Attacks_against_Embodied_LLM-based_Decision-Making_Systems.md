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

该研究提出了首个针对基于大语言模型(LLM)的具身AI决策系统的后门攻击框架(BALD)，系统性地探索了攻击面和触发机制。研究者设计了三种攻击方式：词注入、场景操控和知识注入，针对LLM决策流程的不同环节。实验在自动驾驶和家用机器人任务中(GPT-3.5/LLaMA2/PaLM2)验证了攻击有效性，如车辆加速撞向障碍物、机器人将刀具放床上等。词/知识注入攻击成功率近100%，场景操控攻击成功率65%-90%，且部分攻击无需系统入侵。研究揭示了具身LLM系统的严重安全隐患，亟需加强安全防护。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T23:01:54Z
- **目录日期**: 2025-05-01
