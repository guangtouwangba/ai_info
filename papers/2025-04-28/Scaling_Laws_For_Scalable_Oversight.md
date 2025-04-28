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

该研究提出了一个量化监督成功概率的框架，将可扩展监督建模为能力不匹配玩家间的博弈。玩家拥有基于通用智力的监督和欺骗Elo评分，并存在任务不熟练和饱和两个平台期。研究通过改良Nim游戏验证框架，并应用于四种监督博弈场景。分析发现当监督对象比监督者强400Elo时，嵌套可扩展监督(NSO)成功率低于52%，且随能力差距扩大进一步下降。研究还推导了NSO成功条件和最优监督层级数，为控制超级智能系统提供了量化评估方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-28T15:01:07Z
- **目录日期**: 2025-04-28
