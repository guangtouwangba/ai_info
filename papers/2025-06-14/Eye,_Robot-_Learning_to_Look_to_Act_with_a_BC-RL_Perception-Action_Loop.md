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

研究人员开发了EyeRobot系统，通过强化学习训练机械眼球主动注视以辅助任务完成。该系统采用360°摄像头收集演示数据，在模拟环境中训练手眼协同策略：手部（模仿学习）基于眼球观测执行动作，而眼球（强化学习）通过优化手部动作预测获得奖励。其仿生中央凹架构实现了高分辨率低算力需求，带来更稳定的注视点、目标跟踪及抗干扰能力。实验表明，EyeRobot在环绕机械臂的5个全景操作任务中，仅用单摄像头即可通过有效的手眼协调实现大范围操作。项目网址：https://www.eyerobot.net/（100字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-14T19:01:30Z
- **目录日期**: 2025-06-14
