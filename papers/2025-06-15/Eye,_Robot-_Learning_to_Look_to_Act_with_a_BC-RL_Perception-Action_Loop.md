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

EyeRobot是一个受人类主动视觉启发的机器人系统，通过机械眼球旋转观察环境，结合强化学习训练注视策略。系统采用360°摄像头收集遥操作数据，并在仿真环境中训练手眼协同策略：手部（模仿学习）基于眼球观测执行动作，眼部（强化学习）通过手部动作准确性获得奖励，从而自然形成任务导向的注视行为。其仿中央凹的架构在低算力下实现高分辨率观测，表现出稳定注视、目标跟踪和抗干扰能力。实验表明，EyeRobot在环绕式全景操作任务中，仅用单摄像头即可高效完成大范围操作。详情见项目网站：https://www.eyerobot.net/（98字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-15T17:01:52Z
- **目录日期**: 2025-06-15
