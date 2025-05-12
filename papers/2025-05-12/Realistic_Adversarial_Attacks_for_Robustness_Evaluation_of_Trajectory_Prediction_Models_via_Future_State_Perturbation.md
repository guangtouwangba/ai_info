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

轨迹预测是自动驾驶系统的关键功能，用于预判其他道路使用者的动向。现有对抗性攻击方法通常仅扰动周围车辆的过去轨迹，可能导致不现实的测试场景并忽略关键漏洞。本研究提出通过同时扰动对抗车辆的过去和未来状态，结合动态约束和战术行为保持，生成更真实有效的攻击轨迹。新方法在先进预测模型上的测试显示，对抗条件下预测误差和碰撞率显著上升，并暴露出模型无法识别潜在碰撞等弱点。结果表明需要更全面的对抗测试来提升自动驾驶轨迹预测模型的可靠性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T05:01:34Z
- **目录日期**: 2025-05-12
