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

研究团队提出了首个城市空中视觉目标搜索（AVOS）基准数据集CityAVOS，包含6类2420个任务，用于评估无人机自主搜索能力。同时开发了PRPSearcher方法，基于多模态大语言模型（MLLMs），模拟人类三层认知：通过动态语义地图增强空间感知，基于语义吸引力的3D认知地图进行目标推理，以及3D不确定性地图平衡探索与利用。该方法还包含去噪机制和启发式提示策略，实验显示其成功率（+37.69%）和搜索效率显著优于基线，但与人类表现仍有差距，凸显了提升语义推理和空间探索能力的必要性。数据集和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T00:01:30Z
- **目录日期**: 2025-05-15
