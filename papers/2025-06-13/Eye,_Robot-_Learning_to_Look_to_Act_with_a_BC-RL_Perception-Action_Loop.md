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

研究人员开发了EyeRobot系统，通过强化学习训练机械眼球自主注视策略，实现手眼协同完成任务。该系统采用360°摄像头采集数据，在仿真环境中训练眼球注视策略，并通过BC-RL联合训练方法：手部(BC)基于眼球观测学习动作，眼球(RL)则在手部动作正确时获得奖励。这种仿人眼中央凹的架构能以较小计算成本实现高分辨率观测，产生更稳定的注视行为，提升目标跟踪和抗干扰能力。在5个全景工作空间操作任务测试中，EyeRobot仅用单个摄像头就实现了大范围工作空间的有效操作。项目网站：https://www.eyerobot.net/ (100字)

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T05:01:20Z
- **目录日期**: 2025-06-13
