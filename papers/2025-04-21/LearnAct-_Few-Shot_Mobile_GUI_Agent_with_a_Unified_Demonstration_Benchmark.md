# LearnAct: Few-Shot Mobile GUI Agent with a Unified Demonstration Benchmark

**URL**: http://arxiv.org/abs/2504.13805v1

## 原始摘要

Mobile GUI agents show promise in automating tasks but face generalization
challenges in diverse real-world scenarios. Traditional approaches using
pre-training or fine-tuning with massive datasets struggle with the diversity
of mobile applications and user-specific tasks. We propose enhancing mobile GUI
agent capabilities through human demonstrations, focusing on improving
performance in unseen scenarios rather than pursuing universal generalization
through larger datasets. To realize this paradigm, we introduce LearnGUI, the
first comprehensive dataset specifically designed for studying
demonstration-based learning in mobile GUI agents, comprising 2,252 offline
tasks and 101 online tasks with high-quality human demonstrations. We further
develop LearnAct, a sophisticated multi-agent framework that automatically
extracts knowledge from demonstrations to enhance task completion. This
framework integrates three specialized agents: DemoParser for knowledge
extraction, KnowSeeker for relevant knowledge retrieval, and ActExecutor for
demonstration-enhanced task execution. Our experimental results show
significant performance gains in both offline and online evaluations. In
offline assessments, a single demonstration improves model performance,
increasing Gemini-1.5-Pro's accuracy from 19.3% to 51.7%. In online
evaluations, our framework enhances UI-TARS-7B-SFT's task success rate from
18.1% to 32.8%. LearnAct framework and LearnGUI benchmark establish
demonstration-based learning as a promising direction for more adaptable,
personalized, and deployable mobile GUI agents.


## AI 摘要

本文提出LearnGUI数据集和LearnAct框架，通过人类示范提升移动GUI代理在未知场景下的任务执行能力。LearnGUI包含2,252个离线任务和101个在线任务的高质量示范数据。LearnAct框架包含DemoParser、KnowSeeker和ActExecutor三个智能体，分别负责知识提取、检索和执行。实验表明：单次示范使Gemini-1.5-Pro准确率从19.3%提升至51.7%；在线评估中，UI-TARS-7B-SFT任务成功率从18.1%提升至32.8%。该研究为开发更具适应性、个性化的移动GUI代理提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T09:01:43Z
- **目录日期**: 2025-04-21
