# Towards Autonomous UAV Visual Object Search in City Space: Benchmark and Agentic Methodology

**URL**: http://arxiv.org/abs/2505.08765v1

## 原始摘要

Aerial Visual Object Search (AVOS) tasks in urban environments require
Unmanned Aerial Vehicles (UAVs) to autonomously search for and identify target
objects using visual and textual cues without external guidance. Existing
approaches struggle in complex urban environments due to redundant semantic
processing, similar object distinction, and the exploration-exploitation
dilemma. To bridge this gap and support the AVOS task, we introduce CityAVOS,
the first benchmark dataset for autonomous search of common urban objects. This
dataset comprises 2,420 tasks across six object categories with varying
difficulty levels, enabling comprehensive evaluation of UAV agents' search
capabilities. To solve the AVOS tasks, we also propose PRPSearcher
(Perception-Reasoning-Planning Searcher), a novel agentic method powered by
multi-modal large language models (MLLMs) that mimics human three-tier
cognition. Specifically, PRPSearcher constructs three specialized maps: an
object-centric dynamic semantic map enhancing spatial perception, a 3D
cognitive map based on semantic attraction values for target reasoning, and a
3D uncertainty map for balanced exploration-exploitation search. Also, our
approach incorporates a denoising mechanism to mitigate interference from
similar objects and utilizes an Inspiration Promote Thought (IPT) prompting
mechanism for adaptive action planning. Experimental results on CityAVOS
demonstrate that PRPSearcher surpasses existing baselines in both success rate
and search efficiency (on average: +37.69% SR, +28.96% SPL, -30.69% MSS, and
-46.40% NE). While promising, the performance gap compared to humans highlights
the need for better semantic reasoning and spatial exploration capabilities in
AVOS tasks. This work establishes a foundation for future advances in embodied
target search. Dataset and source code are available at
https://anonymous.4open.science/r/CityAVOS-3DF8.


## AI 摘要

本文提出了首个城市空中视觉目标搜索(CityAVOS)基准数据集，包含6类共2420个任务，并开发了基于多模态大语言模型的PRPSearcher方法。该方法通过构建动态语义地图、3D认知地图和不确定性地图，模拟人类三层认知机制，结合去噪和启发式提示技术，显著提升了无人机在复杂城市场景中的目标搜索能力。实验表明，PRPSearcher在成功率(+37.69%)和搜索效率(+28.96% SPL)等指标上优于现有基线，但与人类表现仍存在差距，凸显了提升语义推理和空间探索能力的必要性。该研究为具身目标搜索领域奠定了基础。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T04:01:50Z
- **目录日期**: 2025-05-14
