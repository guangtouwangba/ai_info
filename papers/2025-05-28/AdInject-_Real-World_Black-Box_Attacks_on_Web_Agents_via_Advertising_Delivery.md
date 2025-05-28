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

基于视觉语言模型（VLM）的Web代理在模拟人类与网站交互方面取得进展，但面临安全漏洞。现有攻击方法常依赖不现实的假设（如直接修改HTML或获取用户意图）。本文提出AdInject，一种更贴近现实的攻击方法，通过互联网广告投放注入恶意内容，无需了解代理参数或用户意图。AdInject利用VLM优化广告内容，使其看似相关，诱导代理点击。实验显示攻击成功率超60%，部分场景接近100%，揭示了广告投放作为攻击媒介的风险。研究强调需开发防御机制应对此类威胁。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T14:01:20Z
- **目录日期**: 2025-05-28
