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

研究人员开发了EyeRobot系统，这是一个受人类主动视觉行为启发的机器人系统。该系统通过机械眼球自由旋转观察环境，并使用强化学习训练注视策略。研究团队首先收集了360度摄像头与遥操作演示数据，并在仿真环境中训练手-眼协同策略：手部（行为克隆）从眼球渲染的观察中学习，而眼部（强化学习）在手部正确预测动作时获得奖励。该系统采用类似中央凹的高分辨率策略架构，能稳定注视并追踪目标。实验表明，EyeRobot能在单摄像头下高效完成大范围操作任务。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-15T16:01:25Z
- **目录日期**: 2025-06-15
