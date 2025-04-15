# SeaView: Software Engineering Agent Visual Interface for Enhanced Workflow

**URL**: http://arxiv.org/abs/2504.08696v1

## 原始摘要

Auto-regressive LLM-based software engineering (SWE) agents, henceforth SWE
agents, have made tremendous progress (&gt;60% on SWE-Bench Verified) on
real-world coding challenges including GitHub issue resolution. SWE agents use
a combination of reasoning, environment interaction and self-reflection to
resolve issues thereby generating "trajectories". Analysis of SWE agent
trajectories is difficult, not only as they exceed LLM sequence length
(sometimes, greater than 128k) but also because it involves a relatively
prolonged interaction between an LLM and the environment managed by the agent.
In case of an agent error, it can be hard to decipher, locate and understand
its scope. Similarly, it can be hard to track improvements or regression over
multiple runs or experiments. While a lot of research has gone into making
these SWE agents reach state-of-the-art, much less focus has been put into
creating tools to help analyze and visualize agent output. We propose a novel
tool called SeaView: Software Engineering Agent Visual Interface for Enhanced
Workflow, with a vision to assist SWE-agent researchers to visualize and
inspect their experiments. SeaView's novel mechanisms help compare experimental
runs with varying hyper-parameters or LLMs, and quickly get an understanding of
LLM or environment related problems. Based on our user study, experienced
researchers spend between 10 and 30 minutes to gather the information provided
by SeaView, while researchers with little experience can spend between 30
minutes to 1 hour to diagnose their experiment.


## AI 摘要

基于自回归大语言模型（LLM）的软件工程（SWE）智能体在解决实际编码问题（如GitHub问题修复）方面取得显著进展（SWE-Bench Verified准确率>60%），但其生成的复杂交互轨迹（可能超过128k长度）难以分析。现有研究多聚焦于提升智能体性能，而缺乏分析工具。为此，研究者提出可视化工具SeaView，帮助研究人员快速比较不同超参数或LLM的实验结果，定位LLM或环境相关问题。用户研究表明，SeaView可将资深研究员的诊断时间从10-30分钟、新手的30-60分钟大幅缩短，显著提升实验分析效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T00:01:43Z
- **目录日期**: 2025-04-15
