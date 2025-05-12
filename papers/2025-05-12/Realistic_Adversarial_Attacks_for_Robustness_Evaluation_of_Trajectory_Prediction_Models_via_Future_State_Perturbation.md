# Realistic Adversarial Attacks for Robustness Evaluation of Trajectory Prediction Models via Future State Perturbation

**URL**: http://arxiv.org/abs/2505.06134v1

## 原始摘要

Trajectory prediction is a key element of autonomous vehicle systems,
enabling them to anticipate and react to the movements of other road users.
Evaluating the robustness of prediction models against adversarial attacks is
essential to ensure their reliability in real-world traffic. However, current
approaches tend to focus on perturbing the past positions of surrounding
agents, which can generate unrealistic scenarios and overlook critical
vulnerabilities. This limitation may result in overly optimistic assessments of
model performance in real-world conditions.
  In this work, we demonstrate that perturbing not just past but also future
states of adversarial agents can uncover previously undetected weaknesses and
thereby provide a more rigorous evaluation of model robustness. Our novel
approach incorporates dynamic constraints and preserves tactical behaviors,
enabling more effective and realistic adversarial attacks. We introduce new
performance measures to assess the realism and impact of these adversarial
trajectories. Testing our method on a state-of-the-art prediction model
revealed significant increases in prediction errors and collision rates under
adversarial conditions. Qualitative analysis further showed that our attacks
can expose critical weaknesses, such as the inability of the model to detect
potential collisions in what appear to be safe predictions. These results
underscore the need for more comprehensive adversarial testing to better
evaluate and improve the reliability of trajectory prediction models for
autonomous vehicles.


## AI 摘要

本文提出了一种更全面的对抗性测试方法，通过扰动自动驾驶轨迹预测模型中周围车辆的历史和未来状态，揭示传统方法可能忽略的关键弱点。该方法结合动态约束并保持战术行为，生成更真实有效的对抗轨迹。实验表明，该方法能显著增加预测误差和碰撞率，暴露模型在看似安全预测中无法检测潜在碰撞等缺陷。研究强调需要更严格的对抗测试来评估和提高自动驾驶轨迹预测模型的可靠性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T14:01:41Z
- **目录日期**: 2025-05-12
