# Scaling Laws For Scalable Oversight

**URL**: http://arxiv.org/abs/2504.18530v1

## 原始摘要

Scalable oversight, the process by which weaker AI systems supervise stronger
ones, has been proposed as a key strategy to control future superintelligent
systems. However, it is still unclear how scalable oversight itself scales. To
address this gap, we propose a framework that quantifies the probability of
successful oversight as a function of the capabilities of the overseer and the
system being overseen. Specifically, our framework models oversight as a game
between capability-mismatched players; the players have oversight-specific and
deception-specific Elo scores that are a piecewise-linear function of their
general intelligence, with two plateaus corresponding to task incompetence and
task saturation. We validate our framework with a modified version of the game
Nim and then apply it to four oversight games: "Mafia", "Debate", "Backdoor
Code" and "Wargames". For each game, we find scaling laws that approximate how
domain performance depends on general AI system capability (using Chatbot Arena
Elo as a proxy for general capability). We then build on our findings in a
theoretical study of Nested Scalable Oversight (NSO), a process in which
trusted models oversee untrusted stronger models, which then become the trusted
models in the next step. We identify conditions under which NSO succeeds and
derive numerically (and in some cases analytically) the optimal number of
oversight levels to maximize the probability of oversight success. In our
numerical examples, the NSO success rate is below 52% when overseeing systems
that are 400 Elo points stronger than the baseline overseer, and it declines
further for overseeing even stronger systems.


## AI 摘要

本文提出一个量化可扩展监督（scalable oversight）成功概率的框架，将监督建模为能力不匹配玩家间的博弈，双方监督和欺骗能力随通用智能呈分段线性变化。研究通过改良Nim游戏验证框架，并应用于"黑手党""辩论""后门代码""战争游戏"四种监督场景，发现领域性能与AI通用能力（以Chatbot Arena Elo为指标）存在缩放规律。进一步理论分析嵌套式可扩展监督（NSO）表明：当被监督系统比基线监督者强400 Elo时，NSO成功率低于52%，且随能力差距扩大持续下降。研究还推导了最优监督层级数。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T08:10:55Z
- **目录日期**: 2025-04-28
