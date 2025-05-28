# AdInject: Real-World Black-Box Attacks on Web Agents via Advertising Delivery

**URL**: http://arxiv.org/abs/2505.21499v1

## 原始摘要

Vision-Language Model (VLM) based Web Agents represent a significant step
towards automating complex tasks by simulating human-like interaction with
websites. However, their deployment in uncontrolled web environments introduces
significant security vulnerabilities. Existing research on adversarial
environmental injection attacks often relies on unrealistic assumptions, such
as direct HTML manipulation, knowledge of user intent, or access to agent model
parameters, limiting their practical applicability. In this paper, we propose
AdInject, a novel and real-world black-box attack method that leverages the
internet advertising delivery to inject malicious content into the Web Agent's
environment. AdInject operates under a significantly more realistic threat
model than prior work, assuming a black-box agent, static malicious content
constraints, and no specific knowledge of user intent. AdInject includes
strategies for designing malicious ad content aimed at misleading agents into
clicking, and a VLM-based ad content optimization technique that infers
potential user intents from the target website's context and integrates these
intents into the ad content to make it appear more relevant or critical to the
agent's task, thus enhancing attack effectiveness. Experimental evaluations
demonstrate the effectiveness of AdInject, attack success rates exceeding 60%
in most scenarios and approaching 100% in certain cases. This strongly
demonstrates that prevalent advertising delivery constitutes a potent and
real-world vector for environment injection attacks against Web Agents. This
work highlights a critical vulnerability in Web Agent security arising from
real-world environment manipulation channels, underscoring the urgent need for
developing robust defense mechanisms against such threats. Our code is
available at https://github.com/NicerWang/AdInject.


## AI 摘要

本文提出AdInject，一种基于互联网广告投放的新型黑盒攻击方法，针对视觉语言模型(VLM)驱动的网页代理。该方法通过恶意广告注入误导代理点击，无需知晓用户意图或代理参数，攻击成功率高达60%-100%。相比现有依赖HTML操纵或模型参数的方法，AdInject在更现实的威胁模型下运作，利用网站上下文优化广告内容以增强攻击性。研究揭示了广告投放渠道对网页代理安全的严重威胁，亟需开发防御机制。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T08:01:25Z
- **目录日期**: 2025-05-28
