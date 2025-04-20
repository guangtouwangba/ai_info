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

在差分隐私的序贯学习环境中，智能体通过内生噪声保护隐私。针对连续信号，研究提出了一种基于阈值距离的自适应平滑随机响应机制，相比传统均匀噪声方法，能更有效融合私有信号和历史观察，将渐近学习速度提升至$\Theta_{\epsilon}(\log(n))$（隐私预算$\epsilon>0$时），显著优于非隐私场景的$\Theta(\sqrt{\log(n)})$。有限隐私预算还能避免早期智能体长期错误决策的问题。然而，对于二元信号，隐私噪声会阻碍学习，因智能体在信息级联形成前倾向于使用恒定随机响应策略，降低了动作信息量。研究揭示了隐私机制对连续/离散信号学习效率的差异化影响。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-20T01:30:25Z
- **目录日期**: 2025-04-20
