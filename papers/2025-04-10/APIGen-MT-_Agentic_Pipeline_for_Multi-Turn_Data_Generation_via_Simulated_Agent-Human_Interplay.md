# APIGen-MT: Agentic Pipeline for Multi-Turn Data Generation via Simulated Agent-Human Interplay

**URL**: http://arxiv.org/abs/2504.03601v2

## 原始摘要

Training effective AI agents for multi-turn interactions requires
high-quality data that captures realistic human-agent dynamics, yet such data
is scarce and expensive to collect manually. We introduce APIGen-MT, a
two-phase framework that generates verifiable and diverse multi-turn agent
data. In the first phase, our agentic pipeline produces detailed task
blueprints with ground-truth actions, leveraging a committee of LLM reviewers
and iterative feedback loops. These blueprints are then transformed into
complete interaction trajectories through simulated human-agent interplay. We
train a family of models -- the xLAM-2-fc-r series with sizes ranging from 1B
to 70B parameters. Our models outperform frontier models such as GPT-4o and
Claude 3.5 on $\tau$-bench and BFCL benchmarks, with the smaller models
surpassing their larger counterparts, particularly in multi-turn settings,
while maintaining superior consistency across multiple trials. Comprehensive
experiments demonstrate that our verified blueprint-to-details approach yields
high-quality training data, enabling the development of more reliable,
efficient, and capable agents. We open-source both the synthetic data collected
and the trained xLAM-2-fc-r models to advance research in AI agents. Models are
available on HuggingFace at
https://huggingface.co/collections/Salesforce/xlam-2-67ef5be12949d8dcdae354c4
and project website is https://apigen-mt.github.io


## AI 摘要

APIGen-MT是一个两阶段框架，用于生成可验证且多样化的多轮AI代理训练数据。第一阶段通过LLM评审委员会和迭代反馈生成详细任务蓝图，第二阶段模拟人机交互生成完整轨迹。基于该数据训练的xLAM-2-fc-r模型系列（1B-70B参数）在τ-bench和BFCL基准测试中超越GPT-4o和Claude 3.5，其中小模型在多轮交互中表现尤为突出且更稳定。研究证实该方法能产生高质量训练数据，提升代理的可靠性、效率和能力。相关数据和模型已在HuggingFace开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T00:02:43Z
- **目录日期**: 2025-04-10
