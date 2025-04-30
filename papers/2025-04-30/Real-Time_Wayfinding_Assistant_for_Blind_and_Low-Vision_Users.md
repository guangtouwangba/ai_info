# Real-Time Wayfinding Assistant for Blind and Low-Vision Users

**URL**: http://arxiv.org/abs/2504.20976v1

## 原始摘要

Navigating unfamiliar places continues to be one of the most persistent and
essential everyday obstacles for those who are blind or have limited vision
(BLV). Existing assistive technologies, such as GPS-based navigation systems,
AI-powered smart glasses, and sonar-equipped canes, often face limitations in
real-time obstacle avoidance, precise localization, and adaptability to dynamic
surroundings. To investigate potential solutions, we introduced PathFinder, a
novel map-less navigation system that explores different models for
understanding 2D images, including Vision Language Models (VLMs), Large
Language Models (LLMs), and employs monocular depth estimation for free-path
detection. Our approach integrates a Depth-First Search (DFS) algorithm on
depth images to determine the longest obstacle-free path, ensuring optimal
route selection while maintaining computational efficiency. We conducted
comparative evaluations against existing AI-powered navigation methods and
performed a usability study with BLV participants. The results demonstrate that
PathFinder achieves a favorable balance between accuracy, computational
efficiency, and real-time responsiveness. Notably, it reduces mean absolute
error (MAE) and improves decision-making speed in outdoor navigation compared
to AI-based alternatives. Participant feedback emphasizes the system's
usability and effectiveness in outside situations, but also identifies issues
in complicated indoor locations and low-light conditions. Usability testing
revealed that 73% of participants understood how to use the app in about a
minute, and 80% praised its balance of accuracy, quick response, and overall
convenience.


## AI 摘要

PathFinder是一种新型无地图导航系统，专为视障人士设计，通过视觉语言模型(VLM)、大语言模型(LLM)和单目深度估计实现实时路径规划。该系统采用深度优先搜索(DFS)算法分析深度图像，找出最长无障碍路径，兼顾计算效率和路线优化。测试显示，PathFinder在户外导航中显著降低平均绝对误差(MAE)，决策速度优于现有AI方案。73%的视障用户能在1分钟内掌握使用，80%认可其准确性、响应速度和便利性，但在复杂室内环境和弱光条件下仍需改进。研究突显了该系统在实时避障和环境适应性方面的优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-30T23:01:31Z
- **目录日期**: 2025-04-30
