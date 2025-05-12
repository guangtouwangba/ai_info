# Scaling Laws For Scalable Oversight

**URL**: http://arxiv.org/abs/2504.18530v2

## 原始摘要

Scalable oversight, the process by which weaker AI systems supervise stronger
ones, has been proposed as a key strategy to control future superintelligent
systems. However, it is still unclear how scalable oversight itself scales. To
address this gap, we propose a framework that quantifies the probability of
successful oversight as a function of the capabilities of the overseer and the
system being overseen. Specifically, our framework models oversight as a game
between capability-mismatched players; the players have oversight-specific Elo
scores that are a piecewise-linear function of their general intelligence, with
two plateaus corresponding to task incompetence and task saturation. We
validate our framework with a modified version of the game Nim and then apply
it to four oversight games: Mafia, Debate, Backdoor Code and Wargames. For each
game, we find scaling laws that approximate how domain performance depends on
general AI system capability. We then build on our findings in a theoretical
study of Nested Scalable Oversight (NSO), a process in which trusted models
oversee untrusted stronger models, which then become the trusted models in the
next step. We identify conditions under which NSO succeeds and derive
numerically (and in some cases analytically) the optimal number of oversight
levels to maximize the probability of oversight success. We also apply our
theory to our four oversight games, where we find that NSO success rates at a
general Elo gap of 400 are 13.5% for Mafia, 51.7% for Debate, 10.0% for
Backdoor Code, and 9.4% for Wargames; these rates decline further when
overseeing stronger systems.


## AI 摘要

该研究提出了一个量化监督成功概率的框架，将可扩展监督建模为能力不匹配玩家间的博弈。监督者和被监督者拥有基于通用智力的分段线性Elo评分，包含"任务无能"和"任务饱和"两个平台期。研究通过改良Nim游戏验证框架，并应用于黑手党、辩论、后门代码和战争游戏四种监督场景，发现领域表现与AI能力间的缩放规律。针对嵌套可扩展监督(NSO)的理论分析表明，在400通用Elo差距下，四种场景的监督成功率分别为13.5%、51.7%、10.0%和9.4%，且随被监督系统变强而进一步降低。研究还推导了最优监督层级数以最大化成功率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-12T18:02:26Z
- **目录日期**: 2025-05-12
