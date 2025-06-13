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

EyeRobot是一个受人类主动视觉启发的机器人系统，通过机械眼球旋转观察环境并完成现实任务。研究团队使用强化学习训练凝视策略，首先通过360°摄像头收集遥操作演示数据，并在仿真环境中渲染任意眼球视角。采用BC-RL联合训练框架：手部（BC）从眼球观测学习动作，眼部（RL）则在手部预测正确时获得奖励，从而自然形成手眼协调——眼球会主动看向有助于任务完成的区域。该系统采用类似中央凹的高效视觉架构，在五个全景操作任务中表现出稳定的注视、目标跟踪和抗干扰能力，仅用单摄像头即可实现大范围操作。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T11:01:10Z
- **目录日期**: 2025-06-13
