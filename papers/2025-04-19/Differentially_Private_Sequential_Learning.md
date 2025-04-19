# Differentially Private Sequential Learning

**URL**: http://arxiv.org/abs/2502.19525v3

## 原始摘要

In a differentially private sequential learning setting, agents introduce
endogenous noise into their actions to maintain privacy. Applying this to a
standard sequential learning model leads to different outcomes for continuous
vs. binary signals. For continuous signals with a nonzero privacy budget, we
introduce a novel smoothed randomized response mechanism that adapts noise
based on distance to a threshold, unlike traditional randomized response, which
applies uniform noise. This enables agents' actions to better reflect both
private signals and observed history, accelerating asymptotic learning speed to
$\Theta_{\epsilon}(\log(n))$, compared to $\Theta(\sqrt{\log(n)})$ in the
non-private regime where privacy budget is infinite. Moreover, in the
non-private setting, the expected stopping time for the first correct decision
and the number of incorrect actions diverge, meaning early agents may make
mistakes for an unreasonably long period. In contrast, under a finite privacy
budget $\epsilon \in (0,1)$, both remain finite, highlighting a stark contrast
between private and non-private learning. Learning with continuous signals in
the private regime is more efficient, as smooth randomized response enhances
the log-likelihood ratio over time, improving information aggregation.
Conversely, for binary signals, differential privacy noise hinders learning, as
agents tend to use a constant randomized response strategy before an
information cascade forms, reducing action informativeness and hampering the
overall process.


## AI 摘要

在差分隐私的序贯学习中，代理通过内生噪声保护隐私。针对连续信号，本文提出基于阈值距离的自适应平滑随机响应机制（传统方法采用均匀噪声），使行动更有效结合私有信号与历史观察，将渐近学习速度提升至Θₑ(log n)（非隐私场景为Θ(√log n)）。有限隐私预算ε∈(0,1)下，首个正确决策的预期停止时间和错误行动次数均保持有限，而传统场景二者发散。连续信号下隐私学习更高效，平滑机制随时间增强对数似然比；但二进制信号中隐私噪声会阻碍学习，因代理在信息级联前倾向使用恒定随机响应策略，降低行动信息量。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-19T10:02:17Z
- **目录日期**: 2025-04-19
