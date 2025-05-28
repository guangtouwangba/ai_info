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

基于视觉语言模型(VLM)的Web代理在自动化网页交互方面取得进展，但面临严重安全风险。研究者提出AdInject攻击方法，通过互联网广告投放向代理环境注入恶意内容。该方法采用更现实的威胁模型：不依赖HTML操控、不知用户意图、不访问代理参数。通过设计诱导点击的广告内容，并利用VLM优化广告与网页情境的相关性，攻击成功率超过60%，某些情况下接近100%。研究表明广告投放是Web代理的实际攻击渠道，凸显了开发防御机制的紧迫性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T04:03:09Z
- **目录日期**: 2025-05-28
