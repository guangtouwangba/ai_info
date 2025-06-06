# A Smooth Sea Never Made a Skilled $\texttt{SAILOR}$: Robust Imitation via Learning to Search

**URL**: http://arxiv.org/abs/2506.05294v1

## 原始摘要

The fundamental limitation of the behavioral cloning (BC) approach to
imitation learning is that it only teaches an agent what the expert did at
states the expert visited. This means that when a BC agent makes a mistake
which takes them out of the support of the demonstrations, they often don't
know how to recover from it. In this sense, BC is akin to giving the agent the
fish -- giving them dense supervision across a narrow set of states -- rather
than teaching them to fish: to be able to reason independently about achieving
the expert's outcome even when faced with unseen situations at test-time. In
response, we explore learning to search (L2S) from expert demonstrations, i.e.
learning the components required to, at test time, plan to match expert
outcomes, even after making a mistake. These include (1) a world model and (2)
a reward model. We carefully ablate the set of algorithmic and design decisions
required to combine these and other components for stable and
sample/interaction-efficient learning of recovery behavior without additional
human corrections. Across a dozen visual manipulation tasks from three
benchmarks, our approach $\texttt{SAILOR}$ consistently out-performs
state-of-the-art Diffusion Policies trained via BC on the same data.
Furthermore, scaling up the amount of demonstrations used for BC by
5-10$\times$ still leaves a performance gap. We find that $\texttt{SAILOR}$ can
identify nuanced failures and is robust to reward hacking. Our code is
available at https://github.com/arnavkj1995/SAILOR .


## AI 摘要

行为克隆（BC）方法在模仿学习中的根本局限是，它仅能教会智能体在专家访问的状态下模仿专家行为。一旦智能体犯错并偏离演示数据范围，往往无法自主恢复。为此，研究者提出基于专家演示的"学习搜索"（L2S）方法，通过训练世界模型和奖励模型，使智能体能在测试时规划恢复路径。新方法SAILOR在12项视觉操作任务中均优于当前最先进的扩散策略（Diffusion Policies），即使BC使用5-10倍演示数据仍存在性能差距。SAILOR能识别细微错误并抵抗奖励作弊，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T16:02:09Z
- **目录日期**: 2025-06-06
