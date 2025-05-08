# Optimal Deterministic Rendezvous in Labeled Lines

**URL**: http://arxiv.org/abs/2505.04564v1

## 原始摘要

In a rendezvous task, a set of mobile agents dispersed in a network have to
gather at an arbitrary common site. We consider the rendezvous problem on the
infinite labeled line, with $2$ initially asleep agents, without communication,
and a synchronous notion of time. Nodes are labeled with unique positive
integers. The initial distance between the two agents is denoted by $D$. Time
is divided into rounds. We count time from when an agent first wakes up, and
denote by $\tau$ the delay between the agents' wake up times. If awake in a
given round $T$, an agent has three options: stay at its current node $v$, take
port $0$, or take port $1$. If it decides to stay, the agent is still at node
$v$ in round $T+1$. Otherwise, it is at one of the two neighbors of $v$ on the
line, based on the port it chose. The agents achieve rendezvous in $T$ rounds
if they are at the same node in round $T$. We aim for a deterministic algorithm
for this task.
  The problem was recently considered by Miller and Pelc [DISC 2023]. With
$\ell_{\max}$ the largest label of the two starting nodes, they showed that no
algorithm can guarantee rendezvous in $o(D \log^* \ell_{\max})$ rounds. The
lower bound follows from a connection with the LOCAL model of distributed
computing, and holds even if the agents are guaranteed simultaneous wake-up
($\tau = 0$) and are given $D$ as advice. Miller and Pelc also gave an
algorithm of optimal matching complexity $O(D \log^* \ell_{\max})$ when $D$ is
known to the agents, but only obtained the higher bound of $O(D^2 (\log^*
\ell_{\max})^3)$ when $D$ is unknown.
  We improve this second complexity to a tight $O(D \log^* \ell_{\max})$. In
fact, our algorithm achieves rendezvous in $O(D \log^* \ell_{\min})$ rounds,
where $\ell_{\min}$ is the smallest label within distance $O(D)$ of the two
starting positions.


## AI 摘要

本文研究了在无限标记线上两个异步唤醒的移动代理的确定性和会问题。代理初始距离为D，节点具有唯一正整数标签。Miller和Pelc此前证明了最优时间复杂度下界为Ω(D log* ℓ_max)，并给出了当D已知时的最优算法O(D log* ℓ_max)，但对未知D的情况只得到O(D²(log* ℓ_max)³)的结果。本文改进了未知D时的算法，达到了与已知D时相同的最优复杂度O(D log* ℓ_max)，实际上可在O(D log* ℓ_min)轮内完成和会，其中ℓ_min是起始位置O(D)距离内的最小标签。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-08T09:02:36Z
- **目录日期**: 2025-05-08
