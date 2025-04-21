# Parameter-Efficient Continual Fine-Tuning: A Survey

**URL**: http://arxiv.org/abs/2504.13822v1

## 原始摘要

The emergence of large pre-trained networks has revolutionized the AI field,
unlocking new possibilities and achieving unprecedented performance. However,
these models inherit a fundamental limitation from traditional Machine Learning
approaches: their strong dependence on the \textit{i.i.d.} assumption hinders
their adaptability to dynamic learning scenarios. We believe the next
breakthrough in AI lies in enabling efficient adaptation to evolving
environments -- such as the real world -- where new data and tasks arrive
sequentially. This challenge defines the field of Continual Learning (CL), a
Machine Learning paradigm focused on developing lifelong learning neural
models. One alternative to efficiently adapt these large-scale models is known
Parameter-Efficient Fine-Tuning (PEFT). These methods tackle the issue of
adapting the model to a particular data or scenario by performing small and
efficient modifications, achieving similar performance to full fine-tuning.
However, these techniques still lack the ability to adjust the model to
multiple tasks continually, as they suffer from the issue of Catastrophic
Forgetting. In this survey, we first provide an overview of CL algorithms and
PEFT methods before reviewing the state-of-the-art on Parameter-Efficient
Continual Fine-Tuning (PECFT). We examine various approaches, discuss
evaluation metrics, and explore potential future research directions. Our goal
is to highlight the synergy between CL and Parameter-Efficient Fine-Tuning,
guide researchers in this field, and pave the way for novel future research
directions.


## AI 摘要

当前AI领域的大型预训练模型虽取得突破性进展，但仍受限于传统机器学习对独立同分布(i.i.d.)假设的依赖，难以适应动态学习环境。持续学习(CL)致力于开发终身学习的神经网络模型，而参数高效微调(PEFT)方法通过少量参数调整实现高效适配。然而现有PEFT技术仍面临灾难性遗忘问题，无法持续适应多任务场景。本文综述了CL算法与PEFT方法的融合研究(PECFT)，分析现有技术路径、评估指标及未来方向，旨在促进CL与参数高效微调的协同发展，为新型终身学习算法研究提供指引。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-21T05:01:35Z
- **目录日期**: 2025-04-21
