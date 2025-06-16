# Self-Regulating Cars: Automating Traffic Control in Free Flow Road Networks

**URL**: http://arxiv.org/abs/2506.11973v1

## 原始摘要

Free-flow road networks, such as suburban highways, are increasingly
experiencing traffic congestion due to growing commuter inflow and limited
infrastructure. Traditional control mechanisms, such as traffic signals or
local heuristics, are ineffective or infeasible in these high-speed,
signal-free environments. We introduce self-regulating cars, a reinforcement
learning-based traffic control protocol that dynamically modulates vehicle
speeds to optimize throughput and prevent congestion, without requiring new
physical infrastructure. Our approach integrates classical traffic flow theory,
gap acceptance models, and microscopic simulation into a physics-informed RL
framework. By abstracting roads into super-segments, the agent captures
emergent flow dynamics and learns robust speed modulation policies from
instantaneous traffic observations. Evaluated in the high-fidelity PTV Vissim
simulator on a real-world highway network, our method improves total throughput
by 5%, reduces average delay by 13%, and decreases total stops by 3% compared
to the no-control setting. It also achieves smoother, congestion-resistant flow
while generalizing across varied traffic patterns, demonstrating its potential
for scalable, ML-driven traffic management.


## AI 摘要

针对郊区高速公路等自由流路网日益严重的拥堵问题，研究者提出了一种基于强化学习的自调节车辆控制方案。该方法通过动态调节车速来优化通行效率，无需新增物理设施。结合经典交通流理论和微观仿真，系统将道路抽象为超级路段，实时学习最优调速策略。在PTV Vissim仿真平台上对真实高速公路网络的测试表明：相比无控制状态，该方案使总通行量提升5%，平均延误降低13%，停车次数减少3%，并能适应不同交通模式，实现更平稳、抗拥堵的车流。这展现了机器学习在可扩展交通管理中的应用潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T03:26:07Z
- **目录日期**: 2025-06-16
