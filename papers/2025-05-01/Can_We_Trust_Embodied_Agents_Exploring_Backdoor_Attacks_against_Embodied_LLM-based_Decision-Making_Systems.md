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

大型语言模型(LLMs)在具身AI决策任务中展现出潜力，但其微调过程存在严重安全隐患。本研究首次提出针对LLM决策系统的后门攻击框架(BALD)，开发了三种攻击方式：词语注入、场景操纵和知识注入，针对决策流程不同环节。实验在自动驾驶和家庭机器人任务中(GPT-3.5/LLaMA2/PaLM2模型)验证了攻击有效性，如车辆加速撞向障碍物或机器人将刀具放床上。词语和知识注入攻击成功率近100%，场景操纵攻击成功率65-90%，且多数无需运行时入侵。研究揭示了具身LLM系统的关键安全漏洞，亟需防护措施。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T18:02:07Z
- **目录日期**: 2025-05-01
