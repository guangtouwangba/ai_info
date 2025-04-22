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

该研究提出通过人类演示增强移动GUI代理的能力，而非依赖大规模数据集。团队开发了首个演示学习专用数据集LearnGUI（含2,252离线任务和101在线任务），并构建多智能体框架LearnAct，包含演示解析器、知识检索器和执行器三个模块。实验显示：单次演示使Gemini-1.5-Pro准确率从19.3%提升至51.7%；在线测试中UI-TARS-7B-SFT任务成功率从18.1%提高到32.8%。该成果为开发更具适应性和个性化的移动GUI代理提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T01:29:05Z
- **目录日期**: 2025-04-22
