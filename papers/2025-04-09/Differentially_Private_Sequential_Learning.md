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

这篇论文研究了差分隐私下的序贯学习模型，比较了连续信号和二元信号的不同表现。对于连续信号，作者提出了一种基于阈值距离的自适应平滑随机响应机制，相比传统均匀噪声方法，能更好地结合私有信号和历史观察，将渐近学习速度提升至Θε(log(n))，优于非隐私场景的Θ(√log(n))。有限隐私预算还能避免早期代理长期犯错的问题。然而，二元信号在隐私保护下学习效率降低，因为代理在信息级联形成前倾向于使用固定随机响应策略，削弱了行动信息量。研究揭示了隐私学习与非隐私学习的显著差异。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-09T05:02:10Z
- **目录日期**: 2025-04-09
