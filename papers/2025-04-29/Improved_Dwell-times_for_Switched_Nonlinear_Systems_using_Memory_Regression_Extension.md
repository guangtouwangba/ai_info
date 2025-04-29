# Improved Dwell-times for Switched Nonlinear Systems using Memory Regression Extension

**URL**: http://arxiv.org/abs/2504.18457v1

## 原始摘要

This paper presents a switched systems approach for extending the dwell-time
of an autonomous agent during GPS-denied operation by leveraging memory
regressor extension (MRE) techniques. To maintain accurate trajectory tracking
despite unknown dynamics and environmental disturbances, the agent periodically
acquires access to GPS, allowing it to correct accumulated state estimation
errors. The motivation for this work arises from the limitations of existing
switched system approaches, where increasing estimation errors during
GPS-denied intervals and overly conservative dwell-time conditions restrict the
operational efficiency of the agent. By leveraging MRE techniques during
GPS-available intervals, the developed method refines the estimates of unknown
system parameters, thereby enabling longer and more reliable operation in
GPS-denied environments. A Lyapunov-based switched-system stability analysis
establishes that improved parameter estimates obtained through concurrent
learning allow extended operation in GPS-denied intervals without compromising
closed-loop system stability. Simulation results validate the theoretical
findings, demonstrating dwell-time extensions and enhanced trajectory tracking
performance.


## AI 摘要

本文提出了一种基于切换系统的方法，通过利用记忆回归扩展(MRE)技术来延长自主代理在GPS失效环境中的运行时间。该方法通过周期性获取GPS信号来校正累积的状态估计误差，从而在未知动态和环境干扰下保持精确的轨迹跟踪。相比现有切换系统方法，MRE技术改进了系统参数估计，允许更长的GPS失效运行时间而不影响闭环系统稳定性。李雅普诺夫稳定性分析证明了该方法的有效性，仿真结果验证了理论发现，展示了运行时间的延长和轨迹跟踪性能的提升。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-29T02:29:26Z
- **目录日期**: 2025-04-29
