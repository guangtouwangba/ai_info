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

针对郊区高速等无信号自由流路网日益严重的拥堵问题，本研究提出了一种基于强化学习的自调节车辆控制协议。该方法融合交通流理论、间隙接受模型和微观仿真，构建物理信息驱动的RL框架，通过将道路抽象为超级路段来捕捉车流动态，仅需瞬时交通观测即可实时优化车速调节策略。在PTV Vissim仿真平台上对真实高速公路网络的测试表明：相比无控制场景，该方案使总通行量提升5%，平均延误降低13%，停车次数减少3%，并能形成更平稳、抗拥堵的车流，展现出机器学习驱动交通管理的扩展潜力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-16T02:33:58Z
- **目录日期**: 2025-06-16
