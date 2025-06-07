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

行为克隆（BC）方法在模仿学习中的核心局限是只能模仿专家在特定状态下的行为，一旦偏离专家示范的状态，代理难以自主恢复。为解决这一问题，研究者提出从专家演示中学习搜索（L2S）的方法，通过训练世界模型和奖励模型，使代理能在测试时规划恢复策略，无需额外人工修正。新方法SAILOR在12项视觉操作任务中均优于当前最佳的扩散策略（BC），即使BC数据量增加5-10倍仍存在性能差距。SAILOR能识别细微错误并避免奖励作弊，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T19:01:42Z
- **目录日期**: 2025-06-07
