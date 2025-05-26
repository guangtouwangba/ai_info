# One RL to See Them All: Visual Triple Unified Reinforcement Learning

**URL**: http://arxiv.org/abs/2505.18129v1

## 原始摘要

Reinforcement learning (RL) has significantly advanced the reasoning
capabilities of vision-language models (VLMs). However, the use of RL beyond
reasoning tasks remains largely unexplored, especially for perceptionintensive
tasks like object detection and grounding. We propose V-Triune, a Visual Triple
Unified Reinforcement Learning system that enables VLMs to jointly learn visual
reasoning and perception tasks within a single training pipeline. V-Triune
comprises triple complementary components: Sample-Level Data Formatting (to
unify diverse task inputs), Verifier-Level Reward Computation (to deliver
custom rewards via specialized verifiers) , and Source-Level Metric Monitoring
(to diagnose problems at the data-source level). We further introduce a novel
Dynamic IoU reward, which provides adaptive, progressive, and definite feedback
for perception tasks handled by V-Triune. Our approach is instantiated within
off-the-shelf RL training framework using open-source 7B and 32B backbone
models. The resulting model, dubbed Orsta (One RL to See Them All),
demonstrates consistent improvements across both reasoning and perception
tasks. This broad capability is significantly shaped by its training on a
diverse dataset, constructed around four representative visual reasoning tasks
(Math, Puzzle, Chart, and Science) and four visual perception tasks (Grounding,
Detection, Counting, and OCR). Subsequently, Orsta achieves substantial gains
on MEGA-Bench Core, with improvements ranging from +2.1 to an impressive +14.1
across its various 7B and 32B model variants, with performance benefits
extending to a wide range of downstream tasks. These results highlight the
effectiveness and scalability of our unified RL approach for VLMs. The V-Triune
system, along with the Orsta models, is publicly available at
https://github.com/MiniMax-AI.


## AI 摘要

V-Triune提出了一种视觉三重统一强化学习系统，使视觉语言模型(VLM)能通过单一训练流程同时学习视觉推理和感知任务。该系统包含三个核心组件：样本级数据格式化统一任务输入、验证器级奖励计算提供定制化反馈、源级指标监控诊断数据问题，并创新性地提出动态IoU奖励机制。基于开源7B/32B骨干模型实例化的Orsta模型，在数学/谜题/图表/科学等4类推理任务和定位/检测/计数/OCR等4类感知任务上表现全面提升，在MEGA-Bench Core基准上获得+2.1至+14.1的显著提升，证明了统一RL框架的有效性和扩展性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T16:01:34Z
- **目录日期**: 2025-05-26
