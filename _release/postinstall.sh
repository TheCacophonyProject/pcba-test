#!/bin/bash
echo "enable service"
systemctl daemon-reload
systemctl enable pcba-test-interface
