# Disillusion: An Open-Source Deception Defense Product Based on Generative Models

Disillusion (aka "Illusion-Breaking Spell") is a deception defense tool designed for Linux hosts, especially for container usage scenarios. Utilizing the simulation capabilities of generative models, it brings a new form of "honeybait" to hosts - command honeybait.

Honeybait is a widely adopted defense method in deception defense. Usually in various file formats, it uses sensitive information to attract attackers' attention, guiding them into honeynets using file information; or when the honeybait file is opened, it can trigger relevant alerting capabilities for defenders.

The advantages and disadvantages of honeybait are extremely obvious!

The most attractive advantage is that it generates almost no additional performance overhead except for occupying a small amount of storage space.

The disappointing aspect is that it's too passive, especially when attackers raise their vigilance, making it difficult for them to fall for the honeybait.

In the era of large models, we came up with the idea of developing Disillusion. Using the simulation capabilities of generative models, a single file can provide excellent deception defense effects for attacker information gathering scenarios.

## Principle

In Linux system environment variables, there are usually multiple configuration items pointing to executable file directories. By placing Disillusion in directories with "higher execution priority", we can achieve the purpose of "hijacking" common command executions.

After hijacking commands, on one hand, it alerts the defenders, and **on the other hand, it uses the simulation capabilities of large models to mislead, confuse, and delay the attacker's attack.**

The usage method only requires a few simple steps:

1. Configure the config file
2. Copy disillusion to a high-priority directory
3. Modify the filename of disillusion
4. Wait...

Video demonstration:

https://v.qq.com/x/page/x3536kwwhr3.html

