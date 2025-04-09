# Differentially Private Sequential Learning

**URL**: http://arxiv.org/abs/2502.19525v2

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

在差分隐私的序贯学习中，代理通过内生噪声保护隐私。连续信号采用新型平滑随机响应机制，根据与阈值的距离自适应添加噪声，相比传统均匀噪声方法，加速了渐近学习速度至Θϵ(log(n))，优于非隐私场景的Θ(√log(n))。有限隐私预算ϵ∈(0,1)下，首个正确决策的停止时间和错误动作次数保持有限，避免了非隐私场景中的长期错误。连续信号下隐私学习更高效，平滑机制随时间增强对数似然比，促进信息聚合。而二元信号中隐私噪声阻碍学习，代理在信息级联前倾向于恒定随机响应策略，降低了动作信息量。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T11:02:04Z
- **目录日期**: 2025-04-09
