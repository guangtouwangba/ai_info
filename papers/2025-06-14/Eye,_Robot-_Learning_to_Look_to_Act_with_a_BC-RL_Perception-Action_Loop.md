# Eye, Robot: Learning to Look to Act with a BC-RL Perception-Action Loop

**URL**: http://arxiv.org/abs/2506.10968v1

## 原始摘要

Humans do not passively observe the visual world -- we actively look in order
to act. Motivated by this principle, we introduce EyeRobot, a robotic system
with gaze behavior that emerges from the need to complete real-world tasks. We
develop a mechanical eyeball that can freely rotate to observe its surroundings
and train a gaze policy to control it using reinforcement learning. We
accomplish this by first collecting teleoperated demonstrations paired with a
360 camera. This data is imported into a simulation environment that supports
rendering arbitrary eyeball viewpoints, allowing episode rollouts of eye gaze
on top of robot demonstrations. We then introduce a BC-RL loop to train the
hand and eye jointly: the hand (BC) agent is trained from rendered eye
observations, and the eye (RL) agent is rewarded when the hand produces correct
action predictions. In this way, hand-eye coordination emerges as the eye looks
towards regions which allow the hand to complete the task. EyeRobot implements
a foveal-inspired policy architecture allowing high resolution with a small
compute budget, which we find also leads to the emergence of more stable
fixation as well as improved ability to track objects and ignore distractors.
We evaluate EyeRobot on five panoramic workspace manipulation tasks requiring
manipulation in an arc surrounding the robot arm. Our experiments suggest
EyeRobot exhibits hand-eye coordination behaviors which effectively facilitate
manipulation over large workspaces with a single camera. See project site for
videos: https://www.eyerobot.net/


## AI 摘要

EyeRobot是一个受人类主动视觉启发的机器人系统，通过机械眼球旋转观察环境并执行任务。研究团队使用强化学习训练凝视策略，首先采集360°摄像头遥操作数据，在仿真环境中模拟眼球视角。采用"行为克隆-强化学习"联合训练机制：手部（BC）基于眼球观测学习动作，眼球（RL）通过手部动作准确性获得奖励，从而自然形成手眼协调——眼球会主动注视对任务完成关键的区域。该系统采用类似中央凹的高效视觉架构，在五个全景操作任务中展现出稳定的注视能力、目标追踪能力和抗干扰性，仅用单摄像头即可实现大范围操作。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-14T02:30:08Z
- **目录日期**: 2025-06-14
